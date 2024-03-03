const {apphost, log} = portal

class Service {

    constructor() {
        this.name = "srv"
    }

    async get(arg) {
        return {
            arg: arg,
            val: "Hello RPC"
        }
    }

    async sum(a, b) {
        return a + b
    }
}

apphost.bindRpcService(Service).catch(log)
