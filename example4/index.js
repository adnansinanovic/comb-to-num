function main() {
	console.log(":: The bit flags")
	const set = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"];
	const subset = ["C", "F", "H"];

	runExample(set, subset);
}

function runExample(set, subset) {
	const indices = findIndices(set, subset);
	console.log(`subset: ${subset}, indices: ${indices}`);

	const encoded = encode(indices);
	console.log(`encoded: ${encoded}`);

	const decodedIndices = decode(encoded);
	console.log(`decodedIndices: ${decodedIndices}, decoded: ${values(set, decodedIndices)}`);
}

function encode(indices) {
    let result  = 0
    for (let i = 0; i < indices.length; i++) {
        const el = indices[i];
        result += Math.pow(2, el);
    }

	return result
}

function decode(encoded) {
	const max = 32; 
    const result = [];
    for (let i = 0; i < max; i++) {

        const p = 1 << i;
        if ((p & encoded) == p) {
            result.push(i);
        }
    }

	return result
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

main();