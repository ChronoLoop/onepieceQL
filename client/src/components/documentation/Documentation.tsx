import { For, ParentComponent, createSignal } from 'solid-js';
import DocSidebar from './DocSidebar';
import MainContentWrapper from '../MainContentWrapper';

export const DocStrong: ParentComponent = (props) => {
    return <strong class="font-bold">{props.children}</strong>;
};

export const DocCode: ParentComponent = (props) => {
    return (
        <code class="px-2 bg-gray-600/10 text-sky-700 dark:text-sky-600 whitespace-normal inline">
            {props.children}
        </code>
    );
};

export const DocParagraph: ParentComponent = (props) => {
    return <p class="text-lg mb-2 ">{props.children}</p>;
};

export const DocH2: ParentComponent<{ id?: string }> = (props) => {
    return (
        <h2
            class="text-2xl font-bold relative mt-5 mb-3 scroll-mt-16"
            id={props.id}
        >
            {props.children}
        </h2>
    );
};

export const DocH3: ParentComponent<{ id?: string }> = (props) => {
    return (
        <h3
            class="text-xl font-bold relative mt-4 mb-2 scroll-mt-16"
            id={props.id}
        >
            {props.children}
        </h3>
    );
};

export type DocTableDataHeader = {
    key: string;
};

export type DocTableDataBody = {
    key: string;
    type: string;
    description: string;
};

type DocTableProps = {
    headers: DocTableDataHeader[];
    body: DocTableDataBody[];
};

export const DocTable = (props: DocTableProps) => {
    return (
        <table class="w-full text-lg mb-10">
            <thead class="text-left">
                <tr class="border-b border-solid border-gray-300 dark:border-slate-700">
                    <For each={props.headers}>
                        {(header) => (
                            <th class="first:pl-0 p-4 font-medium">
                                {header.key}
                            </th>
                        )}
                    </For>
                </tr>
            </thead>
            <tbody class="text-left">
                <For each={props.body}>
                    {(body) => (
                        <tr class="border-b border-solid border-gray-300 dark:border-slate-700">
                            <td class="first:pl-0 p-4">{body.key}</td>
                            <td class="first:pl-0 p-4">{body.type}</td>
                            <td class="first:pl-0 p-4">{body.description}</td>
                        </tr>
                    )}
                </For>
            </tbody>
        </table>
    );
};

export const DocumentationWrapper: ParentComponent = (props) => {
    const [isOpen, setIsOpen] = createSignal(true);

    return (
        <>
            <DocSidebar isOpen={isOpen()} setIsOpen={setIsOpen} />
            <MainContentWrapper sidebarOpen={isOpen()}>
                {props.children}
            </MainContentWrapper>
        </>
    );
};
