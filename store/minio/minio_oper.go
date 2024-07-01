package minio

import (
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"time"
)

func Put(ctx context.Context, bucket string, object string, reader io.Reader, size int64) (string, error) {
	exists, err := cli.BucketExists(ctx, bucket)
	if err != nil {
		return "", err
	}
	if exists {
		info, err := cli.PutObject(ctx, bucket, object, reader, size, minio.PutObjectOptions{})
		if err != nil {
			return "", err
		}
		return info.ETag, nil
	}
	err = cli.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
	if err != nil {
		return "", err
	}

	info, err := cli.PutObject(ctx, bucket, object, reader, size, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return info.ETag, nil
}
func Get(ctx context.Context, bucket string, object string) (string, string, int64, error) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment")
	urls, err := cli.PresignedGetObject(ctx, bucket, object, expire*time.Hour*24, reqParams)
	if err != nil {
		return "", "", 0, err
	}
	info, err := cli.StatObject(ctx, bucket, object, minio.StatObjectOptions{})
	if err != nil {
		return "", "", 0, err
	}

	return info.ETag, urls.String(), info.Size, nil
}
func Remove(ctx context.Context, bucket string, object string, hash string) error {
	exists, err := cli.BucketExists(ctx, bucket)
	if err != nil {
		return err
	}
	if exists {
		info, err := cli.StatObject(ctx, bucket, object, minio.StatObjectOptions{})
		if err != nil {
			return err
		}
		if info.ETag == hash {
			return cli.RemoveObject(ctx, bucket, object, minio.RemoveObjectOptions{
				ForceDelete: true,
			})
		} else {
			return errors.New("md5 not match,internal error")
		}
	} else {
		return errors.New("bucket  not  exist")
	}
}
