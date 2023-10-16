export default function Priority({ priority } : { priority : number }) {
    switch (priority) {
        case 1:
            return <p>very low</p>;
        case 2:
            return <p>low</p>;
        case 3:
            return <p>neutral</p>;
        case 4:
            return <p>high</p>;
        case 5:
            return <p>very high</p>;
        default:
            return <p>unknown</p>;
    }
}
