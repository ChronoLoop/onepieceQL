/* @refresh reload */
import { render } from 'solid-js/web';

import './index.css';
import App from './App';
import 'highlight.js/styles/atom-one-dark.css';

const root = document.getElementById('root');

render(() => <App />, root!);
