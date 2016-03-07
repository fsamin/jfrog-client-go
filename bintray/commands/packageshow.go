package commands

import (
	"fmt"
	"github.com/JFrogDev/jfrog-cli-go/bintray/utils"
	"github.com/JFrogDev/jfrog-cli-go/cliutils"
)

func ShowPackage(packageDetails *utils.VersionDetails, bintrayDetails *cliutils.BintrayDetails) {
	if bintrayDetails.User == "" {
		bintrayDetails.User = packageDetails.Subject
	}
	url := bintrayDetails.ApiUrl + "packages/" + packageDetails.Subject + "/" +
		packageDetails.Repo + "/" + packageDetails.Package

	fmt.Println("Getting package: " + packageDetails.Package)
	resp, body, _, _ := cliutils.SendGet(url, nil, true, bintrayDetails.User, bintrayDetails.Key)
	if resp.StatusCode == 200 {
		fmt.Println(cliutils.IndentJson(body))
	} else {
		cliutils.Exit(cliutils.ExitCodeError, "Bintray response: "+resp.Status)
	}
}