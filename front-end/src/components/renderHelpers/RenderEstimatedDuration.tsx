export default function RenderEstimatedDuration({ duration }: { duration: number }) {
        switch (duration) {
            case 1:
                return <p>very short</p>;
            case 2:
                return <p>short</p>;
            case 3:
                return <p>standard</p>;
            case 4:
                return <p>long</p>;
            case 5:
                return <p>very long</p>;
            default:
                return <p>unknown</p>;
        }
    }
    