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
    CHARACTER_SCHEMA_BODY,
    FILTER_CHARACTERS_SAMPLE_QUERY,
    FILTER_CHARACTERS_SAMPLE_RESPONSE,
    GET_ALL_CHARACTERS_SAMPLE_QUERY,
    GET_ALL_CHARACTERS_SAMPLE_RESPONSE,
    GET_SINGLE_CHARACTER_SAMPLE_QUERY,
    GET_SINGLE_CHARACTER_SAMPLE_RESPONSE,
} from './constants/character';
import { DEFAULT_TABLE_SCHEMA_HEADERS } from './constants/table';

const CharacterSchemaDoc = () => {
    return (
        <>
            <DocH3 id="character-schema">Character Schema</DocH3>
            <DocTable
                headers={DEFAULT_TABLE_SCHEMA_HEADERS}
                body={CHARACTER_SCHEMA_BODY}
            />
        </>
    );
};

const GetSingleCharacterDoc = () => {
    return (
        <>
            <DocH3 id="single-character">Get Single Character</DocH3>
            <DocStrong>Sample Query</DocStrong>
            <SyntaxHighlighter
                language="graphql"
                codeStr={GET_SINGLE_CHARACTER_SAMPLE_QUERY}
            />
            <DocStrong>Sample Response</DocStrong>
            <SyntaxHighlighter
                language="json"
                codeStr={GET_SINGLE_CHARACTER_SAMPLE_RESPONSE}
            />
        </>
    );
};

const GetAllCharactersDoc = () => {
    return (
        <>
            <DocH3 id="all-characters">Get All Characters</DocH3>
            <DocParagraph>
                Use the <DocCode>characters</DocCode> query to get all
                characters. See{' '}
                <A href="#filter-characters" class="underline" target="_self">
                    <DocStrong>filtering</DocStrong>
                </A>{' '}
                section below for available parameters.
            </DocParagraph>
            <DocStrong>Sample Query</DocStrong>
            <SyntaxHighlighter
                language="graphql"
                codeStr={GET_ALL_CHARACTERS_SAMPLE_QUERY}
            />
            <DocStrong>Sample Response</DocStrong>
            <SyntaxHighlighter
                language="json"
                codeStr={GET_ALL_CHARACTERS_SAMPLE_RESPONSE}
            />
        </>
    );
};

const FilterCharactersDoc = () => {
    return (
        <>
            <DocH3 id="filter-characters">Filter Characters</DocH3>
            <DocStrong>Sample Query</DocStrong>
            <SyntaxHighlighter
                language="graphql"
                codeStr={FILTER_CHARACTERS_SAMPLE_QUERY}
            />
            <DocStrong>Sample Response</DocStrong>
            <SyntaxHighlighter
                language="json"
                codeStr={FILTER_CHARACTERS_SAMPLE_RESPONSE}
            />
        </>
    );
};

const CharacterDoc = () => {
    return (
        <>
            <DocH2 id="character">Character</DocH2>
            <DocParagraph>
                This category contains all canon characters of the World of One
                Piece.
            </DocParagraph>
            <CharacterSchemaDoc />
            <GetSingleCharacterDoc />
            <GetAllCharactersDoc />
            <FilterCharactersDoc />
        </>
    );
};

export default CharacterDoc;
