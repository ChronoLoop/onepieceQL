import { ParentComponent } from 'solid-js';

export const CardList: ParentComponent = (props) => {
    return <ul class="flex flex-col gap-2 text-base">{props.children}</ul>;
};

export const CardListItem: ParentComponent = (props) => {
    return <li>{props.children}</li>;
};

type CardProps = {
    imageSrc?: string;
    imageAlt?: string;
};

const Card: ParentComponent<CardProps> = (props) => {
    return (
        <div class="p-5 border border-gray-300 border-solid rounded-lg shadow-lg bg-white dark:bg-gray-800 dark:border-gray-700">
            {props.imageSrc && (
                <div class="w-full h-64 overflow-hidden mb-2 rounded-t-lg">
                    <img
                        class="object-contain size-full"
                        src={props.imageSrc}
                        alt={props.imageAlt}
                        loading="lazy"
                    />
                </div>
            )}
            {props.children}
        </div>
    );
};

export default Card;
