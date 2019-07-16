package downloader

import (
	gourl "net/url"

	boshblob "github.com/cloudfoundry/bosh-utils/blobstore"
	boshcrypto "github.com/cloudfoundry/bosh-utils/crypto"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

const blobstoreDownloaderLogTag = "BlobstoreDownloader"

type BlobstoreDownloader struct {
	blobstore boshblob.DigestBlobstore
	logger    boshlog.Logger
}

func NewBlobstoreDownloader(
	blobstore boshblob.DigestBlobstore,
	logger boshlog.Logger,
) BlobstoreDownloader {
	return BlobstoreDownloader{
		blobstore: blobstore,
		logger:    logger,
	}
}

// Download takes URL of format blobstore:///blobId?fingerprint=sha1-value
func (d BlobstoreDownloader) Download(url string) (string, error) {
	parsedURL, err := gourl.Parse(url)
	if err != nil {
		return "", bosherr.WrapErrorf(err, "Parsing url %s", url)
	}

	var fingerprint string

	if fingerprints, found := parsedURL.Query()["fingerprint"]; found {
		if len(fingerprints) > 0 {
			fingerprint = fingerprints[0]
		}
	}

	digest := boshcrypto.NewDigest(boshcrypto.DigestAlgorithmSHA1, fingerprint)
	path, err := d.blobstore.Get(parsedURL.Path, digest)
	if err != nil {
		return "", bosherr.WrapError(err, "Downloading blob")
	}

	d.logger.Debug(blobstoreDownloaderLogTag, "Downloaded %s to %s", url, path)

	return path, nil
}

func (d BlobstoreDownloader) CleanUp(path string) error {
	return d.blobstore.CleanUp(path)
}
