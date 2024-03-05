import { createClient, cacheExchange, fetchExchange } from '@urql/core';

export const client = createClient({
    url: '/graphql',
    exchanges: [cacheExchange, fetchExchange],
});
