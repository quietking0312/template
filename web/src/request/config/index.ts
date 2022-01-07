import { ConfigOptions } from "@/request/config/types";

const config: ConfigOptions = {
    base_url: {
        dev: "http://127.0.0.1:9001/api"
    },

    result_code: 0,

    request_timeout: 5000,

    default_headers: 'application/json'
}

export default config
