type CreateIntersectionObserverArgs = {
    root?: HTMLElement;
    onIntersect: () => void;
    threshold?: number;
    rootMargin?: string;
};

export const createIntersectionObserver = ({
    root,
    onIntersect,
    threshold = 1.0,
    rootMargin = '0px',
}: CreateIntersectionObserverArgs) => {
    const observer = new IntersectionObserver(
        (entries) =>
            entries.forEach((entry) => {
                entry.isIntersecting && onIntersect();
            }),
        {
            root,
            rootMargin,
            threshold,
        }
    );
    return observer;
};
