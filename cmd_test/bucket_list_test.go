//go:build integration

package cmd

import (
	"github.com/qiniu/qshell/v2/cmd_test/test"
	"path/filepath"
	"strings"
	"testing"
)

func TestBucketList(t *testing.T) {
	TestBatchCopy(t)

	result, errs := test.RunCmdWithError("listbucket", test.Bucket, "--prefix", "hello")
	if len(errs) > 0 {
		t.Fatal("error:", errs)
	}

	if !strings.Contains(result, test.Key) {
		t.Fatal("expected key:% but not exist", test.Key)
	}

	return
}

func TestBucketListToFile(t *testing.T) {
	rootPath, err := test.ResultPath()
	if err != nil {
		t.Fatal("get root path error:", err)
		return
	}
	file := filepath.Join(rootPath, test.Bucket+"_listbucket.txt")
	_, errs := test.RunCmdWithError("listbucket", test.Bucket, "--prefix", "hello", "-o", file)
	defer test.RemoveFile(file)

	if len(errs) > 0 {
		t.Fatal("error:", errs)
	}

	if !test.IsFileHasContent(file) {
		t.Fatal("list bucket to file error: file empty")
	}

	return
}

func TestBucketListWithNotExistBucket(t *testing.T) {
	_, err := test.RunCmdWithError("listbucket", "test")
	if !strings.Contains(err, "no such bucket") {
		t.Fail()
	}
}

func TestBucketListNoBucket(t *testing.T) {
	_, err := test.RunCmdWithError("listbucket")
	if !strings.Contains(err, "Bucket can't be empty") {
		t.Fail()
	}
}

func TestBucketListDocument(t *testing.T) {
	test.TestDocument("listbucket", t)
}

func TestBucketList2(t *testing.T) {
	result, errs := test.RunCmdWithError("listbucket2", test.Bucket,
		"--prefix", "hello",
		"--readable",
		"--end", "2023-01-12-00-00-00")
	if len(errs) > 0 {
		t.Fatal("error:", errs)
	}

	if !strings.Contains(result, "hello") {
		t.Fatal("no expected key:% but not exist", test.BucketDomain)
	}

	return
}

func TestBucketList2WithApiV1(t *testing.T) {
	result, errs := test.RunCmdWithError("listbucket2", test.Bucket,
		"--prefix", "hello2.json",
		"--readable",
		"--end", "2023-01-12-00-00-00",
		"-d")
	if len(errs) > 0 {
		t.Fatal("error:", errs)
	}

	if !strings.Contains(result, "list by api v1,") {
		t.Fatal("should list by v1")
	}

	result, errs = test.RunCmdWithError("listbucket2", test.Bucket,
		"--api-version", "v1",
		"--prefix", "hello2.json",
		"--readable",
		"--end", "2023-01-12-00-00-00",
		"-d")
	if len(errs) > 0 {
		t.Fatal("error:", errs)
	}

	if !strings.Contains(result, "list by api v1,") {
		t.Fatal("should list by v1")
	}

	if !strings.Contains(result, "hello") {
		t.Fatal("no expected key:% but not exist", test.BucketDomain)
	}

	return
}

func TestBucketList2WithApiV2(t *testing.T) {
	result, errs := test.RunCmdWithError("listbucket2", test.Bucket,
		"--api-version", "v2",
		"--readable",
		"--prefix", "hello2.json",
		"--end", "2023-01-12-00-00-00",
		"-d")
	if len(errs) > 0 {
		t.Fatal("error:", errs)
	}

	if strings.Contains(result, "list by api v2,") {
		t.Fatal("should list by v1")
	}

	if !strings.Contains(result, "api v2 is deprecated") {
		t.Fatal("api v2 should be deprecated")
	}

	if !strings.Contains(result, "hello2.json") {
		t.Fatal("no expected key:% but not exist", test.BucketDomain)
	}

	if strings.Contains(result, "hello3.json") {
		t.Fatal("hello3.json shouldn't be list")
	}

	return
}

func TestBucketList2ToFile(t *testing.T) {
	defaultContent := "AAAAAAA\n"
	file, err := test.CreateFileWithContent(test.Bucket+"-listbucket2.txt", defaultContent)
	if err != nil {
		t.Fatal("get root path error:", err)
		return
	}
	defer test.RemoveFile(file)

	_, errs := test.RunCmdWithError("listbucket2", test.Bucket, "--prefix", "hello",
		"--file-types", "0,1,2",
		"-o", file,
		"--outfile", file)

	if len(errs) > 0 {
		t.Fatal("error:", errs)
	}

	content := test.FileContent(file)
	if strings.HasPrefix(content, defaultContent) {
		t.Fatal("list bucket to file error: should't append")
	}

	if !test.IsFileHasContent(file) {
		t.Fatal("list bucket to file error: file empty")
	}

	return
}

func TestBucketList2ToFileByAppend(t *testing.T) {
	defaultContent := "AAAAAAA\n"
	file, err := test.CreateFileWithContent(test.Bucket+"-listbucket2.txt", defaultContent)
	if err != nil {
		t.Fatal("get root path error:", err)
		return
	}
	defer test.RemoveFile(file)

	_, errs := test.RunCmdWithError("listbucket2", test.Bucket,
		"--prefix", "hello",
		"-o", file,
		"--append")

	if len(errs) > 0 {
		t.Fatal("error:", errs)
	}

	content := test.FileContent(file)
	if !strings.HasPrefix(content, defaultContent) {
		t.Fatal("list bucket to file append error")
	}

	if !test.IsFileHasContent(file) {
		t.Fatal("list bucket to file error: file empty")
	}

	return
}

func TestBucketList2NoBucket(t *testing.T) {
	_, err := test.RunCmdWithError("listbucket2")
	if !strings.Contains(err, "Bucket can't be empty") {
		t.Fail()
	}
}

func TestBucketList2Document(t *testing.T) {
	test.TestDocument("listbucket2", t)
}
