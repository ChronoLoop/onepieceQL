import { DocumentationWrapper } from '../components/documentation/Documentation';
import IntroductionDoc from '../components/documentation/IntroductionDoc';
import CharacterDoc from '../components/documentation/CharacterDoc';
import DevilFruitDoc from '../components/documentation/DevilFruitDoc';

const Home = () => {
    return (
        <DocumentationWrapper>
            <IntroductionDoc />
            <CharacterDoc />
            <DevilFruitDoc />
        </DocumentationWrapper>
    );
};

export default Home;
