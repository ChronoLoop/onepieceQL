import hljs from 'highlight.js/lib/core';

//TODO: try to dynamic imports for languages when this is resolved: https://github.com/vitejs/vite/issues/14102
import javascript from 'highlight.js/lib/languages/javascript';
import json from 'highlight.js/lib/languages/json';
import graphql from 'highlight.js/lib/languages/graphql';

hljs.registerLanguage('graphql', graphql);
hljs.registerLanguage('javascript', javascript);
hljs.registerLanguage('json', json);

type SyntaxLanguage = 'javascript' | 'graphql' | 'json';

type SyntaxHighlighterProps = {
    language: SyntaxLanguage;
    codeStr: string;
    className?: string;
};
const SyntaxHighlighter = (props: SyntaxHighlighterProps) => {
    const { value } = hljs.highlight(props.codeStr, {
        language: props.language,
    });

    return (
        <>
            <pre class="my-3 overflow-auto">
                <code
                    class={`${props.className || ''} hljs`}
                    innerHTML={value}
                    lang={props.language}
                />
            </pre>
        </>
    );
};

export default SyntaxHighlighter;
