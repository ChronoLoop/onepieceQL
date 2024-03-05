import { ParentComponent } from 'solid-js';

type MainContentWrapperProps = {
    sidebarOpen: boolean;
};

const MainContentWrapper: ParentComponent<MainContentWrapperProps> = (
    props
) => {
    return (
        <article
            class="break-words"
            classList={{ ['lg:ml-60']: props.sidebarOpen }}
        >
            <div class="px-20 max-lg:px-14 pt-4 pb-64 max-w-7xl mx-auto">
                {props.children}
            </div>
        </article>
    );
};

export default MainContentWrapper;
