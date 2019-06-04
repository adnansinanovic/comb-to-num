function main() {
    console.log(":: Odometer effect\n");
    const setA = ["John", "Jack", "Joe"];
    const setB = ["Water", "Fire"];
    const setC = ["Brazil", "UK", "Russia", "Japan", "Morocco", "France"];
    const setD = ["Blue", "Red", "Yellow", "White"];

    const lengths = [setA.length, setB.length, setC.length, setD.length];
    const maxCombinations = calcMax(lengths);

    console.log(`num => decoded => ${"elements".padStart(20, ' ').padEnd(30, ' ')} => ${"encoded".padStart(4, ' ')}`)
    console.log("-".repeat(70));

    for (let i = 0; i < maxCombinations; i++) {
        const n = decode(i, lengths);
        const elements = [n[0], n[1], n[2], n[3]];
        const encoded = encode(elements, lengths);
        console.log(`${i.toString().padStart(3, '')} => ${n} => [${setA[n[0]].padStart(4, ' ')}, ${setB[n[1]].padStart(5, ' ')}, ${setC[n[2]].padStart(7, ' ')}, ${setD[n[3]].padStart(6, ' ')}] => ${encoded.toString().padStart(4, ' ')}`);
    }
}

function decode(code, lengths )  {	
	const weights = calculateWeights(lengths);
	const elements = new Array(lengths.length);
	const sums = new Array(lengths.length);

    const ln = weights.length - 1;
    for (let i = ln; i >= 0; i--) {
        if (i == ln) {
            elements[i] = Math.floor(code / weights[i]);
            sums[i] = elements[i] * weights[i];
            continue
        }

        elements[i] = Math.floor((code - sums[i+1]) / weights[i]);
		sums[i] = sums[i+1] + elements[i]*weights[i];
        
    }

	return elements
}

function encode(elements, lengths) {
	const weights = calculateWeights(lengths);

    let code = 0
    for (let i = 0; i < elements.length; i++) {
        const element = elements[i];
        code += (elements[i] * weights[i]);
    };
	

	return code;
}


function calcMax(lengths)  {
    let max = 1;
    for (let i = 0; i < lengths.length; i++) {
        max *= lengths[i];
        
    }

	return max;
}

function calculateWeights(lengths)  {
    const weights = new Array(lengths.length);
    for (let i = 0; i < lengths.length; i++) {
        if (i == 0) {
            weights[0] = 1;
            continue;
        }
        weights[i] = weights[i-1] * lengths[i-1];
        
    }

	return weights;
}


main()