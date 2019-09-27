import os
from datetime import datetime

import github3

# Connect to GitHub API and push the changes.
github = github3.login(token=os.environ["GITHUB_TOKEN"])
release = github.repository(*os.environ["GITHUB_REPOSITORY"].split("/")).create_release(
    f"draft/{datetime.now().isoformat(timespec='seconds')}",
    name="Versão Preliminar",
    body="Versão Preliminar",
    draft=True,
)
release.upload_asset(
    "application/pdf", os.environ["ASSET_PATH"], open(os.environ["ASSET_PATH"], "rb")
)
