# Cyberark PSM Service Check

This simple service checks the status of the Cyberark PSM service and exposes an HTTP endpoint where it reports the state of the PSM service. It serves as a healthcheck endpoint for load balancers.

If the PSM service is running the PSM server will return **PASS** on port 80 if the PSM service is down it will return **FAIL**.

This is a single executable and does not require IIS unlike Cyberark's implementation which does use IIS to do essentially the same. Another advantage is that the executable is installed in a few seconds and uses a tiny amount of resources to run.

## Installation

*(Optional) Have a look at the source code or even build the executable yourself using Golang on a Windows host instead of downloading the installer.*

Download the svccheck.msi from the [releases page](https://github.com/theneedyguy/cyberark-psm-check/releases) and install it on the PSM servers. You can check if it works by navigating to http://localhost on the PSM server. 