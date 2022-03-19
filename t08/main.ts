import { Type } from "@sinclair/typebox";
import Fastify from "fastify";
(async () => {
    const a = (new URL("http://localhost:3000/hello?name=1"))
    console.log(a.host)
    console.log(a.href)
    console.log(a.searchParams)
    console.log(a.hash)
    // const fastify = Fastify({
    //     ajv: {
    //         customOptions: {
    //             removeAdditional: false,
    //             useDefaults: true,
    //             coerceTypes: false,
    //             allErrors: true,
    //         }
    //     }
    // })
    
    // fastify.route({
    //     method: 'GET',
    //     url: '/hello',
    //     async handler(req, rep) {

    //         console.log(req.raw.url)
    //         console.log(req.query)
    //         return 'world'
    //     },
    //     schema: {
    //         querystring: Type.Object({
    //             name: Type.Optional(Type.String()),
    //         }, { additionalProperties: false }),
    //     }
    // })
    // await fastify.listen(3000,)
})()