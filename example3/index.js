function main() {
	// set of all elements
	const set = ["A", "B", "C", "D", "E", "F", "G", "G", "I", "J", "K"];

	// combination of elements from set
	const combination = ["C", "K", "G"];
	runExample(set, combination);
}

function runExample(set, subset) {
	console.log(":: Combinadics")

	// get indices for combination element
	const indices = findIndices(set, subset);
	console.log(`subset [${subset}], indices: [${indices}]`)

    // sort
    indices.sort((x, y) => x < y);
	console.log(`sorted indices: [${indices}]`);

	const encoded = encode(indices);
	console.log(`encoded: ${encoded}`);

	const decodedIndices = decode(encoded, indices.length, 20);

	console.log(`decoded indices: [${decodedIndices}], decoded: ${values(set, decodedIndices)}\n`);
}

function encode(set) {
    let code = 0;
    for (let i = 0; i < set.length; i++) {
        const n = set[i] - 1;
        const m = set.length - i;
        if (n >= m) {
            code += binomial(n, m);
        }
    }

	return code;
}

function decode(encodedValue, subsetLength, totalElements) {
	const result = [];
	let binom = binomial(totalElements, subsetLength);
    
    for (let index = totalElements; index >= 1; index--) {
        if (binom <= encodedValue) {
            result.push(index + 1);
            encodedValue -= binom;
            binom = Math.floor(binom * subsetLength / (index - subsetLength + 1));
            subsetLength--;
        }

        binom = Math.floor(binom * (index - subsetLength) / index);
        
    }

    return result;
}

function values(set , indices)  {
    const subset = [];
    for (let i = 0; i < indices.length; i++) {
        const element = indices[i];
        subset.push(set[element]);
        
    }

	return subset
}

function findIndices(set, subset) {
    const indices = [];
    for (let i = 0; i < subset.length; i++) {
        const element = subset[i];
        indices.push(index(set, element));
        
    }

	return indices
}

function index(set, element) {
    for (let i = 0; i < set.length; i++) {
        const el = set[i];
        if (element == el) {
            return i;
        }
    }
	throw new Error("element does not belong to the set");
}


function binomial(n, k) {
    var coeff = 1;
    var i;
 
    if (k < 0 || k > n) return 0;
 
    for (i = 0; i < k; i++) {
        coeff = coeff * (n - i) / (i + 1);
    }
 
    return coeff;
}
 

main();