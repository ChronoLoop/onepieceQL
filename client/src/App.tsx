import { ParentComponent } from 'solid-js';
import { Router, Route } from '@solidjs/router';
import Navbar from './components/Navbar';
import Home from './pages/Home';
import Explore from './pages/Explore';

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
                <Route path="/explore" component={Explore} />
            </Router>
        </>
    );
};

export default App;
