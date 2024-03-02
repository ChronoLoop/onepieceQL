import { ComponentProps } from 'solid-js';
import Sidebar, { SidebarSection } from '../Sidebar';

type SidebarItems = ComponentProps<typeof SidebarSection>['sidebarItems'];

const INTRODUCTION_SUBSECTION: SidebarItems = [
    { title: 'Contribute', href: '#contribute' },
    { title: 'Info', href: '#info' },
];

const CHARACTER_SUBSECTION: SidebarItems = [
    { title: 'Character Schema', href: '#character-schema' },
    { title: 'Get Single Character', href: '#single-character' },
    { title: 'Get All Characters', href: '#all-characters' },
    { title: 'Filter Characters', href: '#filter-characters' },
];

const DEVIL_FRUIT_SUBSECTION: SidebarItems = [
    { title: 'Devil Fruit Schema', href: '#devil-fruit-schema' },
    { title: 'Get Single Devil Fruit', href: '#single-devil-fruit' },
    { title: 'Get All Devil Fruits', href: '#all-devil-fruits' },
    { title: 'Filter Devil Fruits', href: '#filter-devil-fruits' },
];

type DocSidebarProps = ComponentProps<typeof Sidebar>;

const DocSidebar = (props: DocSidebarProps) => {
    return (
        <Sidebar setIsOpen={props.setIsOpen} isOpen={props.isOpen}>
            <SidebarSection
                header={{
                    title: 'Introduction',
                    href: '#introduction',
                }}
                sidebarItems={INTRODUCTION_SUBSECTION}
            />

            <SidebarSection
                header={{ title: 'Character', href: '#character' }}
                sidebarItems={CHARACTER_SUBSECTION}
            />
            <SidebarSection
                header={{ title: 'Devil Fruit', href: '#devil-fruit' }}
                sidebarItems={DEVIL_FRUIT_SUBSECTION}
            />
        </Sidebar>
    );
};

export default DocSidebar;
