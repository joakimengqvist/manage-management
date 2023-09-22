import Project from '../components/projects/Project';
import { Row, Col } from 'antd';
import { useSelector } from 'react-redux';
import { State } from '../types/state';
import { hasPrivilege } from '../helpers/hasPrivileges';
import { PRIVILEGES } from '../enums/privileges';

const ProjectDetails: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)
    return (
        <div style={{padding: '12px 8px'}}>
            <Row>
                <Col span={8}>
                {hasPrivilege(userPrivileges, PRIVILEGES.project_read) &&
                    <div style={{padding: '4px'}}>
                        <Project />
                    </div>
                }
                </Col>
                <Col span={16}>
                </Col>
            </Row>
        </div>
    )

}

export default ProjectDetails;