import { A } from '@solidjs/router';
import {
    DocCode,
    DocH2,
    DocH3,
    DocParagraph,
    DocStrong,
    DocTable,
} from './Documentation';
import SyntaxHighlighter from '../SyntaxHighlighter';
import { DEFAULT_TABLE_SCHEMA_HEADERS } from './constants/table';
import {
    INTRODUCTION_INFO_SCHEMA,
    INTRODUCTION_SAMPLE_QUERY,
    INTRODUCTION_SAMPLE_RESPONSE,
} from './constants/introduction';

const ContributeDoc = () => {
    return (
        <>
            <DocH3 id="contribute">Contribute</DocH3>
            <DocParagraph>
                Feel free to fork{' '}
                <A
                    href="https://github.com/ChronoLoop/onepieceQL"
                    target="_blank"
                    class="underline"
                >
                    <DocStrong>here</DocStrong>
                </A>{' '}
                and make improvements.
            </DocParagraph>
        </>
    );
};

const InfoDoc = () => {
    return (
        <>
            <DocH2 id="info">Info</DocH2>
            <DocParagraph>
                Use the <DocCode>Info</DocCode> type to get the metadata you
                need for total count and pagination. All types have an{' '}
                <DocCode>Info</DocCode> type on them.
            </DocParagraph>
            <DocH3>Info Schema</DocH3>
            <DocTable
                headers={DEFAULT_TABLE_SCHEMA_HEADERS}
                body={INTRODUCTION_INFO_SCHEMA}
            />
            <DocStrong>Sample Query</DocStrong>
            <SyntaxHighlighter
                language={'graphql'}
                codeStr={INTRODUCTION_SAMPLE_QUERY}
            />
            <DocStrong>Sample Response</DocStrong>
            <SyntaxHighlighter
                language={'json'}
                codeStr={INTRODUCTION_SAMPLE_RESPONSE}
            />
        </>
    );
};

const IntroductionDoc = () => {
    return (
        <>
            <DocH2 id="introduction">Introduction</DocH2>
            <DocParagraph>
                Welcome to the GraphQL API documentation of the beloved anime,
                One Piece! Use this documentation to help you get the most out
                of the <DocStrong>OnePieceQL API</DocStrong>.
            </DocParagraph>
            <DocParagraph>
                You can query the API at{' '}
                <DocCode>https://onepieceql.up.railway.app/graphql</DocCode>
            </DocParagraph>
            <ContributeDoc />
            <InfoDoc />
        </>
    );
};

export default IntroductionDoc;
