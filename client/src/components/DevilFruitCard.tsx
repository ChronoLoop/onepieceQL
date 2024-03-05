import { DevilFruit } from '../services/fruits';
import { truncateString } from '../utils/string';
import Card, { CardList, CardListItem } from './Card';

type DevilFruitCardProps = {
    devilFruit: DevilFruit;
};

const DevilFruitCard = (props: DevilFruitCardProps) => {
    return (
        <Card
            imageAlt={props.devilFruit.englishName}
            imageSrc={props.devilFruit.avatarSrc}
        >
            <CardList>
                <CardListItem>
                    <strong class="font-bold">Name: </strong>
                    {props.devilFruit.englishName}
                </CardListItem>
                <CardListItem>
                    <strong class="font-bold">Type: </strong>
                    {props.devilFruit.type}
                </CardListItem>
                <CardListItem>
                    <strong class="font-bold">Current Owner: </strong>
                    {props.devilFruit.currentOwner}
                </CardListItem>
                <CardListItem>
                    <strong class="font-bold">Previous Owner: </strong>
                    {props.devilFruit.previousOwner}
                </CardListItem>
                <CardListItem>
                    <strong class="font-bold">Meaning: </strong>
                    {props.devilFruit.meaning}
                </CardListItem>
                <CardListItem>
                    <strong class="font-bold">Description: </strong>
                    {truncateString(props.devilFruit.description, 200)}
                </CardListItem>
            </CardList>
        </Card>
    );
};

export default DevilFruitCard;
