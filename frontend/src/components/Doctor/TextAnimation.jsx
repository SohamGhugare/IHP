import React from 'react'
import TextTransition, { presets } from 'react-text-transition';

const TEXTS = ['Precise', 'Detailed', 'Objective'];

const TextAnimation = () => {
    const [index, setIndex] = React.useState(0);

    React.useEffect(() => {
        const intervalId = setInterval(
            () => setIndex((index) => index + 1),
            1500, // every 3 seconds
        );
        return () => clearTimeout(intervalId);
    }, []);

    return (
        <div className="flex flex-col justify-center flex-1 m-10">
            <TextTransition springConfig={presets.wobbly}>
                <text className="text-8xl font-bold text-green-300">
                    {TEXTS[index % TEXTS.length]}
                </text>

            </TextTransition>
            <text className="text-5xl font-bold text-white mt-3">
                medical analysis, virtually.
            </text>
        </div>
    );
}

export default TextAnimation