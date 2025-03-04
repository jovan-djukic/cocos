# IGVM Build Process

We assume that our current working directory is the root of the cocos repository, both on the host machine and in the VM.

## Building IGVM

To build `igvmmeasure`, run the following command from the root of the cocos repository:

```
make cli
```

## Usage
`igvmmeasure` calculates the launch measurement for an IGVM file and can generate a signed version. It ensures integrity by precomputing the expected launch digest, which can be verified against the attestation report. The tool parses IGVM directives, outputs the measurement as a hex string, or creates a signed file for verification at guest launch.

### Example
We measure an IGVM file using our measure command, run:

```bash
./build/cocos-cli igvmmeasure /path/to/igvm/file
```

The tool will parse the directives in the IGVM file, calculate the launch measurement, and output the computed digest. If successful, it prints the measurement to standard output.

Here is a sample output
```
91c4929bec2d0ecf11a708e09f0a57d7d82208bcba2451564444a4b01c22d047995ca27f9053f86de4e8063e9f810548
```