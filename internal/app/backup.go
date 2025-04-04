package app

import (
	"bufio"
	"bytes"
	"fmt"
	"time"

	"github.com/govindkailas/vault-backup/internal/pkg/s3"
	"github.com/govindkailas/vault-backup/internal/pkg/vault"
)

const (
	TIME_LAYOUT         = "20060102-150405"
	SNAPASHOT_EXTENSION = "snap"
)

func Backup(vConfig *vault.Config, s3Config *s3.Client) error {
	fileName := fmt.Sprintf("backup-%s.%s", time.Now().Format(TIME_LAYOUT), SNAPASHOT_EXTENSION)

	fmt.Println("Starting backup...")

	// create vault client
	vaultClient, err := vault.NewClient(vConfig)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// create s3 client
	s3Client := s3.NewClient(s3Config.AccessKey, s3Config.SecretAccessKey, s3Config.Region, s3Config.Bucket, s3Config.Endpoint, s3Config.FileName)

	// create new buffer writer
	buf := bytes.NewBuffer(nil)
	w := bufio.NewWriter(buf)

	// do a vault backup
	err = vaultClient.Backup(w)
	if err != nil {
		fmt.Println(err)

		return err
	}

	// read from buffer
	r := bytes.NewReader(buf.Bytes())

	// put file with backup in s3 bucket
	err = s3Client.PutObject(r, fileName)
	if err != nil {
		fmt.Println(err)

		return err
	}

	// copy this file to latest backup
	err = s3Client.CopyObject(fileName)
	if err != nil {
		fmt.Println(err)

		return err
	}

	// flush the writer
	w.Flush()

	fmt.Printf("Backup with name '%s' created.\n", fileName)

	return nil
}
