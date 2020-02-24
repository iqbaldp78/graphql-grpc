const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require("@apollo/gateway");

const gateway = new ApolloGateway({
    serviceList: [
        { name: 'radiance', url: 'http://localhost:8080/query' },
        // { name: 'accounts', url: 'http://localhost:4001/query' }
    ],
});

const server = new ApolloServer({
    gateway,

    subscriptions: false,
});

server.listen(3000).then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});