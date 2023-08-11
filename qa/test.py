import os
import subprocess

def get_test_path():
    return os.path.join('.', 'tests', '...')

def get_coverage_analysis_path():
    return os.path.join('qa', 'test_coverage.out')

def run(coverage=False):
    args = ['go', 'test', '-v', get_test_path()]

    if coverage:
        args.extend([
            '-coverprofile',
            get_coverage_analysis_path(),
            '-coverpkg',
            os.path.join('.', '...')
        ])

    subprocess.run(
        args,
        check=True
    )

def main():
    run()

if __name__ == '__main__':
    main()

