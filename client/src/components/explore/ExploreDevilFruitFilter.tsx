import { SetStoreFunction } from 'solid-js/store';
import { DevilFruitArgs } from '../../services/fruits';
import { ExploreSidebarListMenu, ExploreSidebarSection } from '.';
import { SidebarHeader } from '../Sidebar';
import Input from '../Input';

const DEVIL_FRUIT_TYPES = ['Logia', 'Zoan', 'Paramecia'];

type ExploreDevilFruitFilterProps = {
    filter: DevilFruitArgs['filter'];
    setFilter: SetStoreFunction<DevilFruitArgs['filter']>;
};

const ExploreDevilFruitFilter = (props: ExploreDevilFruitFilterProps) => {
    const onItemClick =
        (name: keyof DevilFruitArgs['filter']) => (item: string | undefined) =>
            props.setFilter(name, item);
    return (
        <>
            <ExploreSidebarSection>
                <SidebarHeader>Name</SidebarHeader>
                <Input
                    type="text"
                    placeholder="ex: Mochi Mochi no Mi"
                    name="name"
                    id="name"
                    onInput={(e) => props.setFilter('name', e.target.value)}
                    value={props.filter.name}
                />
            </ExploreSidebarSection>
            <ExploreSidebarListMenu
                header="Type"
                items={DEVIL_FRUIT_TYPES}
                selectedItem={props.filter.type}
                onItemClick={onItemClick('type')}
            />
            <ExploreSidebarSection>
                <SidebarHeader>Previous Owner</SidebarHeader>
                <Input
                    type="text"
                    placeholder="ex: Charlotte Katakuri"
                    name="previousOwner"
                    id="previousOwner"
                    onInput={(e) =>
                        props.setFilter('previousOwner', e.target.value)
                    }
                    value={props.filter.name}
                />
            </ExploreSidebarSection>
            <ExploreSidebarSection>
                <SidebarHeader>Current Owner</SidebarHeader>
                <Input
                    type="text"
                    placeholder="ex: Monkey D. Luffy"
                    name="currentOwner"
                    id="currentOwner"
                    onInput={(e) =>
                        props.setFilter('currentOwner', e.target.value)
                    }
                    value={props.filter.name}
                />
            </ExploreSidebarSection>
        </>
    );
};

export default ExploreDevilFruitFilter;
