import subprocess
import frontmatter
from pathlib import Path
from datetime import datetime

collections: set[str] = {"blog"}
CREATED_KEY = "created"
UPDATED_KEY = "updated"


def get_staged_content() -> list[Path]:
    result = subprocess.run(
        ["git", "diff", "--cached", "--name-only", "--diff-filter=AM"],
        capture_output=True,
        text=True,
        check=True
    )
    files = [Path(line.strip()) for line in result.stdout.splitlines()]
    return [
        f for f in files
        if f.suffix == ".md"
        and f.parts[1] == "content"
        and f.parts[2] in collections
    ]


def normalise_datetime(date: str) -> datetime:
    formats = ["%Y-%m-%d", "%Y-%m-%dT%H:%M:%S%z", "%Y-%m-%dT%H:%M:%S.%f%z"]
    for fmt in formats:
        try:
            return datetime.strptime(date, fmt).replace(microsecond=0).astimezone()
        except ValueError:
            continue


def get_git_root() -> str:
    root = subprocess.run(
        ["git", "rev-parse", "--show-toplevel"],
        capture_output=True,
        text=True,
        check=True
    )
    return Path(root.stdout.strip())


def process_date_frontmatter(article: Path, root: str = get_git_root()) -> None:
    abs_path = root / article
    stats = abs_path.stat()
    post = frontmatter.load(abs_path)

    fm_created = post.get(CREATED_KEY)
    fm_updated = post.get(UPDATED_KEY)

    mtime = datetime.fromtimestamp(
        stats.st_mtime).replace(microsecond=0).astimezone()

    fm_created = normalise_datetime(
        fm_created) if fm_created else mtime

    fm_updated = max(normalise_datetime(
        fm_updated),
        mtime) if fm_updated else mtime

    post[CREATED_KEY] = fm_created.isoformat()
    post[UPDATED_KEY] = fm_updated.isoformat()

    with open(abs_path, "w", encoding="utf-8") as f:
        f.write(frontmatter.dumps(post))


if __name__ == "__main__":
    articles: list[Path] = get_staged_content()

    print(f"Processing Dates for: {articles}")
    for a in articles:
        process_date_frontmatter(a)
    print("Date Processing Complete")
