import { A } from '@solidjs/router';
import SyntaxHighlighter from '../SyntaxHighlighter';
import {
    DocCode,
    DocH2,
    DocH3,
    DocParagraph,
    DocStrong,
    DocTable,
} from './Documentation';
import {
    DEVIL_FRUIT_SCHEMA_BODY,
    FILTER_DEVIL_FRUITS_SAMPLE_QUERY,
    FILTER_DEVIL_FRUITS_SAMPLE_RESPONSE,
    GET_ALL_DEVIL_FRUITS_SAMPLE_QUERY,
    GET_ALL_DEVIL_FRUITS_SAMPLE_RESPONSE,
    GET_SINGLE_DEVIL_FRUIT_SAMPLE_QUERY,
    GET_SINGLE_DEVIL_FRUIT_SAMPLE_RESPONSE,
} from './constants/devilfruit';
import { DEFAULT_TABLE_SCHEMA_HEADERS } from './constants/table';

const DevilFruitSchemaDoc = () => {
    return (
        <>
            <DocH3 id="devil-fruit-schema">Devil Fruit Schema</DocH3>
            <DocTable
                headers={DEFAULT_TABLE_SCHEMA_HEADERS}
                body={DEVIL_FRUIT_SCHEMA_BODY}
            />
        </>
    );
};

const GetSingleDevilFruitDoc = () => {
    return (
        <>
            <DocH3 id="single-devil-fruit">Get Single Devil Fruit</DocH3>
            <DocStrong>Sample Query</DocStrong>
            <SyntaxHighlighter
                language="graphql"
                codeStr={GET_SINGLE_DEVIL_FRUIT_SAMPLE_QUERY}
            />
            <DocStrong>Sample Response</DocStrong>
            <SyntaxHighlighter
                language="json"
                codeStr={GET_SINGLE_DEVIL_FRUIT_SAMPLE_RESPONSE}
            />
        </>
    );
};

const GetAllDevilFruitsDoc = () => {
    return (
        <>
            <DocH3 id="all-devil-fruits">Get All Devil Fruits</DocH3>
            <DocParagraph>
                Use the <DocCode>devilFruits</DocCode> query to get all devil
                fruits. See{' '}
                <A href="#filter-devil-fruits" class="underline" target="_self">
                    <DocStrong>filtering</DocStrong>
                </A>{' '}
                section below for available parameters.
            </DocParagraph>
            <DocStrong>Sample Query</DocStrong>
            <SyntaxHighlighter
                language="graphql"
                codeStr={GET_ALL_DEVIL_FRUITS_SAMPLE_QUERY}
            />
            <DocStrong>Sample Response</DocStrong>
            <SyntaxHighlighter
                language="json"
                codeStr={GET_ALL_DEVIL_FRUITS_SAMPLE_RESPONSE}
            />
        </>
    );
};

const FilterDevilFruitsDoc = () => {
    return (
        <>
            <DocH3 id="filter-devil-fruits">Filter Devil Fruits</DocH3>
            <DocStrong>Sample Query</DocStrong>
            <SyntaxHighlighter
                language="graphql"
                codeStr={FILTER_DEVIL_FRUITS_SAMPLE_QUERY}
            />
            <DocStrong>Sample Response</DocStrong>
            <SyntaxHighlighter
                language="json"
                codeStr={FILTER_DEVIL_FRUITS_SAMPLE_RESPONSE}
            />
        </>
    );
};

const DevilFruitDoc = () => {
    return (
        <>
            <DocH2 id="devil-fruit">Devil Fruit</DocH2>
            <DocParagraph>
                This category contains all devil fruits of the World of One
                Piece.
            </DocParagraph>
            <DevilFruitSchemaDoc />
            <GetSingleDevilFruitDoc />
            <GetAllDevilFruitsDoc />
            <FilterDevilFruitsDoc />
        </>
    );
};

export default DevilFruitDoc;
