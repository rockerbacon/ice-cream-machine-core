import os
import subprocess
import sys

STATUS_SUCCESS = 0
STATUS_FAILED = 1
STATUS_INCOMPLETE_COVERAGE = 2

def get_test_path():
    return os.path.join('.', 'tests', '...')

def get_coverage_analysis_path():
    return os.path.join('qa', 'test_coverage.out')

def list_packages():
    main_packages=[
        'internal',
        'pkg'
    ]

    package_roots = [
        f'rockerbacon/ice-cream-machine-core/{pkg}/...'
        for pkg in main_packages
    ]

    args = ['go', 'list']
    args.extend(package_roots)
    process = subprocess.Popen(
        args,
        stdout=subprocess.PIPE,
        text=True
    )

    all_packages=[]
    while line := process.stdout.readline():
        all_packages.append(line[:-1])

    process.wait(1)

    if process.returncode != 0:
        raise Exception("Error while listing packages")

    return all_packages

def run(coverage=False):
    args = ['go', 'test', '-v', get_test_path()]

    if coverage:
        args.extend([
            '-coverprofile',
            get_coverage_analysis_path(),
            '-coverpkg',
            ','.join(list_packages())
        ])

    execution = subprocess.run(
        args,
        stderr=subprocess.PIPE,
        text=True
    )

    if len(execution.stderr):
        print()
        print('Possible files missing tests:')
        print(execution.stderr)
        return STATUS_INCOMPLETE_COVERAGE

    if execution.returncode != 0:
        return STATUS_FAILED

    return STATUS_SUCCESS

def main():
    status = run()

    if status != STATUS_SUCCESS:
        sys.exit(1)

if __name__ == '__main__':
    main()

