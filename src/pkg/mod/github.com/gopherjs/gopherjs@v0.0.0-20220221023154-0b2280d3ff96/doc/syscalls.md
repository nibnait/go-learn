System Calls
------------

System calls are the bridge between your application and your operating system. They are used whenever you access something outside of your application's memory, for example when you write to the console, when you read or write files or when you access the network. In Go, system calls are mostly used by the `os` package, hence the name. When using GopherJS you need to consider if system calls are available or not.

### Output redirection to console

If system calls are not available in your environment (see below), then a special redirection of `os.Stdout` and `os.Stderr` is applied. It buffers a line until it is terminated by a line break and then prints it via JavaScript's `console.log` to your browser's JavaScript console or your system console. That way, `fmt.Println` etc. work as expected, even if system calls are not available.

### In Browser

The JavaScript environment of a web browser is completely isolated from your operating system to protect your machine. You don't want any web page to read or write files on your disk without your consent. That is why system calls are not and will never be available when running your code in a web browser.

### Node.js on Linux and macOS

GopherJS has support for system calls on Linux and macOS. Before running your code with Node.js, you need to install the system calls module. The module is compatible with Node.js version 10.0.0 (or newer). If you want to use an older version you can opt to not install the module, but then system calls are not available.

Compile and install the module with:

```
cd gopherjs/node-syscall/
npm install
```

You can copy build/Release/syscall.node into you `node_modules` directory and run `node -r syscall` to make sure the module can be loaded successfully.

Alternatively, in _your_ `package.json` you can do something like this:

```
{
  "dependencies": {
    "syscall": "file:path/to/gopherjs/node-syscall"
  }
}
```

Which will make `npm install` in your project capable of building the extension. You may need to set `export NODE_PATH="$(npm root)"` to ensure that node can load modules from any working directory, for example when running `gopherjs test`.

### Node.js on Windows

When running your code with Node.js on Windows, it is theoretically possible to use system calls. To do so, you would need a special Node.js module that provides direct access to system calls. However, since the interface is quite different from the one used on Linux and macOS, the system calls module included in GopherJS currently does not support Windows. Sorry. Get in contact if you feel like you want to change this situation.

### Caveats

Note that even with syscalls enabled in Node.js, some programs may not behave as expected due to the fact that the current implementation blocks other goroutines during a syscall, which can lead to a deadlock in some situations. This is not considered a bug, as it is considered sufficient for most test cases (which is all Node.js should be used for). Get in contact if you feel like you want to change this situation.
