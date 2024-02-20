import { createClient, cacheExchange, fetchExchange } from '@urql/core';
import { ParentComponent, createResource } from 'solid-js';
import { Router, Route } from '@solidjs/router';
import Navbar from './components/Navbar';
import Home from './pages/Home';

const client = createClient({
    url: '/graphql',
    exchanges: [cacheExchange, fetchExchange],
});

const DEVIL_FRUIT_QUERY = `
    query ExampleQuery($page: Int) {
      devilFruits(page: $page) {
        results {
          avatarSrc
          currentOwner
          description
          englishName
          id
          japaneseName
          meaning
          previousOwner
          type
          usageDebut
        }
        info {
          count
          next
          pages
          prev
        }
      }
    }
`;
// const devilFruits = createResource(() =>
//     client
//         .query(DEVIL_FRUIT_QUERY, { page: 1 })
//         .toPromise()
//         .then(({ data }) => console.log(data))
// );

const PageWrapper: ParentComponent = (props) => {
    return (
        <>
            <Navbar />
            <main class="w-full ">{props.children}</main>
        </>
    );
};

const App: ParentComponent = () => {
    return (
        <>
            <Router root={PageWrapper}>
                <Route path="/" component={Home} />
            </Router>
        </>
    );
};

export default App;
