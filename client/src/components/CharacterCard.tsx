import { Character } from '../services/characters';
import { truncateString } from '../utils/string';

type CharacterCardProps = {
    character: Character;
};

const CharacterCard = (props: CharacterCardProps) => {
    return (
        <div class="p-5 border border-gray-300 border-solid rounded-lg shadow-lg bg-white dark:bg-gray-800 dark:border-gray-700">
            <div class="w-full h-48 overflow-hidden mb-2 rounded-t-lg">
                <img
                    class="object-contain size-full"
                    src={props.character.avatarSrc}
                    alt={props.character.englishName}
                    loading="lazy"
                />
            </div>
            <ul class="flex flex-col gap-2 text-base">
                <li>
                    <strong class="font-bold">Name: </strong>
                    {props.character.englishName}
                </li>
                <li>
                    <strong class="font-bold">Origin: </strong>
                    {props.character.origin}
                </li>
                <li>
                    <strong class="font-bold">Birthday: </strong>
                    {props.character.birthday}
                </li>
                <li>
                    <strong class="font-bold">Blood type: </strong>
                    {props.character.bloodType}
                </li>
                <li>
                    <strong class="font-bold">Devil Fruit: </strong>
                    {props.character.devilFruitName}
                </li>
                <li>
                    <strong class="font-bold">Description: </strong>
                    {truncateString(props.character.description, 200)}
                </li>
            </ul>
        </div>
    );
};

export default CharacterCard;
