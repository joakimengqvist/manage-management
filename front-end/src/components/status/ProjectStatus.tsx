import { Badge } from 'antd';

const ProjectStatus = ({status} : { status : string}) => {
    switch (status) {
        case 'completed':
          return <Badge status="success" text="Completed" />;
        case 'ongoing':
          return <Badge status="processing" text="Ongoing" />;
        case 'cancelled':
          return <Badge status="error" text="Cancelled" />;
        case 'delayed':
            return <Badge status="warning" text="Delayed" />;
        case 'not-started':
            return <Badge status="default" text="Not started" />;
        default:
          return <Badge status="default" text={status} />;
    }
}

export default ProjectStatus;