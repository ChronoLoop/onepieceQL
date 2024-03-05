import { Character } from '../services/characters';
import { truncateString } from '../utils/string';
import Card, { CardList, CardListItem } from './Card';

type CharacterCardProps = {
    character: Character;
};

const CharacterCard = (props: CharacterCardProps) => {
    return (
        <Card
            imageAlt={props.character.englishName}
            imageSrc={props.character.avatarSrc}
        >
            <CardList>
                <CardListItem>
                    <strong class="font-bold">Name: </strong>
                    {props.character.englishName}
                </CardListItem>
                <CardListItem>
                    <strong class="font-bold">Origin: </strong>
                    {props.character.origin}
                </CardListItem>
                <CardListItem>
                    <strong class="font-bold">Birthday: </strong>
                    {props.character.birthday}
                </CardListItem>
                <CardListItem>
                    <strong class="font-bold">Blood type: </strong>
                    {props.character.bloodType}
                </CardListItem>
                <CardListItem>
                    <strong class="font-bold">Devil Fruit: </strong>
                    {props.character.devilFruitName}
                </CardListItem>
                <CardListItem>
                    <strong class="font-bold">Description: </strong>
                    {truncateString(props.character.description, 200)}
                </CardListItem>
            </CardList>
        </Card>
    );
};

export default CharacterCard;
