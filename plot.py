import re
import matplotlib.pyplot as plt
import sys

if len(sys.argv) < 2:
    print("Usage: python script.py <filename>")
    sys.exit(1)

file_name = sys.argv[1]
# Read data from the text file
with open(file_name, 'r') as file:
    lines = file.readlines()

# Extract variables and times from each line
variables = []
recursive_times = []
iterative_times = []

for line in lines:
    match = re.match(r'Variables: (\d+), Recursive Avg Time: ([\d.]+) seconds, Iterative Avg Time: ([\d.]+) seconds', line)
    if match:
        variables.append(int(match.group(1)))
        recursive_times.append(float(match.group(2)))
        iterative_times.append(float(match.group(3)))

# Plotting
plt.figure(figsize=(8, 6))
plt.plot(variables, recursive_times, label='Recursive')
plt.plot(variables, iterative_times, label='Iterative')
plt.xlabel('Variable Count')
plt.ylabel('Time (seconds)')
plt.title('Recursive vs Iterative Execution Time ' + file_name)
plt.legend()
plt.grid(True)
plt.show()
