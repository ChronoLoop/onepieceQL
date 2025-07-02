import { ParentComponent } from 'solid-js';

const MainContentWrapper: ParentComponent = (props) => {
    return (
        <article class="break-words flex-1 overflow-y-auto h-full max-h-full">
            <div class="px-20 max-lg:px-14 pt-4 pb-64 max-w-7xl mx-auto">
                {props.children}
            </div>
        </article>
    );
};

export default MainContentWrapper;
