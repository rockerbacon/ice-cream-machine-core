import subprocess
import sys
import test

MIN_COVERAGE = 100

def read_total_coverage(line):
    if not line.startswith('total:'):
        return 0

    percentage_str=''
    i = -3
    while (-i < len(line) and line[i] != '\t'):
        percentage_str += line[i]
        i -= 1

    return float(percentage_str[::-1])

def parse_process_output(process):
    total_coverage = 0
    line=''
    while line := process.stdout.readline():
        print(line)
        if total_coverage == 0:
            total_coverage = read_total_coverage(line)

    return {
        'total': total_coverage
    }

def analyze_coverage_output():
    args = [
        'go', 'tool', 'cover',
        '-func', test.get_coverage_analysis_path()
    ]
    process = subprocess.Popen(
        args,
        stdout=subprocess.PIPE,
        text=True
    )

    results = parse_process_output(process)

    process.wait(1)

    return results

def main():
    test.run(coverage=True)
    coverage = analyze_coverage_output()

    if coverage['total'] < MIN_COVERAGE:
        print()
        print(f'Coverage of {coverage["total"]}% does not fufill minimum requirement of {MIN_COVERAGE}%')
        sys.exit(1)

if __name__ == '__main__':
    main()

