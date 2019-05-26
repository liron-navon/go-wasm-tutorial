// turns our wasm functions with callbacks, into functions that return promises
const promisedWASM = (func) => {
    return (...props) => new Promise((resolve, reject) => {
        func(...props, (error, value) => {
            error ? reject(error) : resolve(value)
        })
    })
}

const run = async () => {
    const go = new Go();
    await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
        go.run(result.instance)
    });

    // we bind the callbacks to an empty object
    const main = await promisedWASM(window["__wasm_main"])({})

    const add = promisedWASM(main.add);
    const sum = await add(5, 5);
    
    console.log('sum', sum)
}
run();