import React from 'react'
import TextTransition, { presets } from 'react-text-transition';

const TEXTS = ['Reports,', 'Diagnosis,', 'Invoices,'];

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
        <div className="flex flex-col items-center justify-center">
            <div className= "flex flex-col justify-center flex-2 m-10">

                <TextTransition springConfig={presets.wobbly}>
                    <text className="text-7xl font-bold text-blue-500 text-center mb-5">
                        Your {TEXTS[index % TEXTS.length]}
                    </text>

                </TextTransition>
                <text className="text-7xl font-bold text-violet-900 text-center">

                    Our Responsibility
                </text>
            </div>
            
        </div>
        //     <div className="flex flex-col justify-center flex-2 m-10">
        //     <text className="text-7xl font-bold text-violet-900 text-center mb-5">
        //         Your Reports, Diagnosis, Invoices

        //     </text>
        //     <text className="text-7xl font-bold text-violet-900 text-center">

        //         Our Responsibility
        //     </text>
        // </div>
    );
}

export default TextAnimation