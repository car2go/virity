package api

import (
	"net/http"
	"testing"

	"github.com/car2go/virity/internal/pluginregistry"
)

func TestNew(T *testing.T) {
	api := apiService{
		mux:     http.NewServeMux(),
		statics: newStaticsServer("/Users/dglinka/src/Go/src/github.com/car2go/virity/internal/monitoring/api/client/dist"),
	}
	api.server = &http.Server{
		Addr:    ":8080",
		Handler: api.mux,
	}

	api.Serve()

	request, err := http.NewRequest("GET", "http://localhost:8080/", nil)

	if err != nil {
		T.Error(err)
		return
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		T.Error(err)
		return
	}

	if response.StatusCode != 200 {
		T.Errorf("Server not reachable. Code: %v", response.StatusCode)
	}

}
func TestPush(t *testing.T) {
	api := apiService{
		mux:     http.NewServeMux(),
		statics: newStaticsServer("/Users/dglinka/src/Go/src/github.com/car2go/virity/internal/monitoring/api/client/dist"),
	}
	api.server = &http.Server{
		Addr:    ":8080",
		Handler: api.mux,
	}

	image := pluginregistry.ImageStack{
		MetaData: pluginregistry.Image{
			Tag:     "debian:latest",
			ImageID: "da653cee0545dfbe3c1864ab3ce782805603356a9cc712acc7b3100d9932fa5e",
		},
		Containers: []pluginregistry.Container{
			pluginregistry.Container{
				Name:     "/test",
				Hostname: "localhost",
				Image:    "debian:latest",
				ID:       "bf51c9974229f0a3790366464fef13e2cdbf0be5b682874f4e78f1538005a800",
				ImageID:  "da653cee0545dfbe3c1864ab3ce782805603356a9cc712acc7b3100d9932fa5e",
			}},
		Vuln: pluginregistry.Vulnerabilities{
			Scanner: "anchore",
			CVE: []pluginregistry.CVE{
				pluginregistry.CVE{
					Fix:      "None",
					Package:  "tar-1.29b-1.1",
					Severity: pluginregistry.SeverityNegligible,
					URL:      "https://security-tracker.debian.org/tracker/CVE-2005-2541",
					Vuln:     "CVE-2005-2541",
				},
				pluginregistry.CVE{
					Fix:      "0.0.9+deb9u1",
					Package:  "sensible-utils-0.0.9",
					Severity: pluginregistry.SeverityMedium,
					URL:      "https://security-tracker.debian.org/tracker/CVE-2017-17512",
					Vuln:     "CVE-2017-17512",
				}, pluginregistry.CVE{
					Fix:      "None",
					Package:  "util-linux-2.29.2-1",
					Severity: pluginregistry.SeverityHigh,
					URL:      "https://security-tracker.debian.org/tracker/CVE-2016-2779",
					Vuln:     "CVE-2016-2779",
				}, pluginregistry.CVE{
					Fix:      "None",
					Package:  "apt-1.4.8",
					Severity: pluginregistry.SeverityLow,
					URL:      "https://security-tracker.debian.org/tracker/CVE-2011-3374",
					Vuln:     "CVE-2011-3374",
				}, pluginregistry.CVE{
					Fix:      "YES",
					Package:  "apt-1.9.8",
					Severity: pluginregistry.SeverityLow,
					URL:      "https://security-tracker.debian.org/tracker/CVE-2011-3374",
					Vuln:     "CVE-2011-3374",
				},
			},
		},
	}

	api.Push(image, pluginregistry.StatusError)

}

func TestResolve(t *testing.T) {
	api := apiService{
		mux:     http.NewServeMux(),
		statics: newStaticsServer("/Users/dglinka/src/Go/src/github.com/car2go/virity/internal/monitoring/api/client/dist"),
	}
	api.server = &http.Server{
		Addr:    ":8080",
		Handler: api.mux,
	}

	image := pluginregistry.ImageStack{
		MetaData: pluginregistry.Image{
			Tag:     "debian:latest",
			ImageID: "da653cee0545dfbe3c1864ab3ce782805603356a9cc712acc7b3100d9932fa5e",
		},
		Containers: []pluginregistry.Container{
			pluginregistry.Container{
				Name:     "/test",
				Hostname: "localhost",
				Image:    "debian:latest",
				ID:       "bf51c9974229f0a3790366464fef13e2cdbf0be5b682874f4e78f1538005a800",
				ImageID:  "da653cee0545dfbe3c1864ab3ce782805603356a9cc712acc7b3100d9932fa5e",
			}},
		Vuln: pluginregistry.Vulnerabilities{
			Scanner: "anchore",
			CVE: []pluginregistry.CVE{
				pluginregistry.CVE{
					Fix:      "None",
					Package:  "tar-1.29b-1.1",
					Severity: pluginregistry.SeverityNegligible,
					URL:      "https://security-tracker.debian.org/tracker/CVE-2005-2541",
					Vuln:     "CVE-2005-2541",
				},
				pluginregistry.CVE{
					Fix:      "0.0.9+deb9u1",
					Package:  "sensible-utils-0.0.9",
					Severity: pluginregistry.SeverityMedium,
					URL:      "https://security-tracker.debian.org/tracker/CVE-2017-17512",
					Vuln:     "CVE-2017-17512",
				}, pluginregistry.CVE{
					Fix:      "None",
					Package:  "util-linux-2.29.2-1",
					Severity: pluginregistry.SeverityHigh,
					URL:      "https://security-tracker.debian.org/tracker/CVE-2016-2779",
					Vuln:     "CVE-2016-2779",
				}, pluginregistry.CVE{
					Fix:      "None",
					Package:  "apt-1.4.8",
					Severity: pluginregistry.SeverityLow,
					URL:      "https://security-tracker.debian.org/tracker/CVE-2011-3374",
					Vuln:     "CVE-2011-3374",
				}, pluginregistry.CVE{
					Fix:      "YES",
					Package:  "apt-1.9.8",
					Severity: pluginregistry.SeverityLow,
					URL:      "https://security-tracker.debian.org/tracker/CVE-2011-3374",
					Vuln:     "CVE-2011-3374",
				},
			},
		},
	}

	api.Resolve(image)
}