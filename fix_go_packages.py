import os
import re
import argparse

PACKAGE_PATTERN = re.compile(r"^package\s+(\w+)", re.MULTILINE)

def fix_package_names(root_dir: str, dry_run: bool = False):
    for dirpath, _, filenames in os.walk(root_dir):
        for filename in filenames:
            if not filename.endswith(".go"):
                continue

            file_path = os.path.join(dirpath, filename)
            folder_name = os.path.basename(dirpath)
            if not folder_name:
                continue

            try:
                with open(file_path, "r", encoding="utf-8") as f:
                    content = f.read()

                match = PACKAGE_PATTERN.search(content)
                if not match:
                    continue

                current_package = match.group(1)
                is_test_file = filename.endswith("_test.go")
                has_test_suffix = current_package.endswith("_test")

                # Logic
                if is_test_file:
                    if has_test_suffix:
                        # Replace prefix but keep _test suffix
                        expected_package = f"{folder_name}_test"
                    else:
                        # No _test suffix → use just folder name
                        expected_package = folder_name
                else:
                    expected_package = folder_name

                if current_package != expected_package:
                    if dry_run:
                        print(f"🔍 [DRY-RUN] Would update {file_path}: {current_package} → {expected_package}")
                    else:
                        new_content = PACKAGE_PATTERN.sub(
                            f"package {expected_package}", content, count=1
                        )
                        with open(file_path, "w", encoding="utf-8") as f:
                            f.write(new_content)
                        print(f"✅ Updated {file_path}: {current_package} → {expected_package}")

            except Exception as e:
                print(f"⚠️ Skipped {file_path}: {e}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Fix Go package declarations to match folder names.")
    parser.add_argument("--path", default=".", help="Root directory (default: current folder)")
    parser.add_argument("--dry-run", action="store_true", help="Preview changes without modifying files")
    args = parser.parse_args()

    fix_package_names(args.path, dry_run=args.dry_run)
