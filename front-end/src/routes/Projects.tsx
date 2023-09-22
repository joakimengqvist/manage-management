import Projects from '../components/projects/Projects';
import CreateProject from '../components/projects/CreateProject';
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
                <Col span={16}>
                {hasPrivilege(userPrivileges, PRIVILEGES.project_read) &&
                    <div style={{padding: '4px'}}>
                        <Projects />
                    </div>
                }   
                </Col>
                <Col span={8}>
                {hasPrivilege(userPrivileges, PRIVILEGES.project_write) &&
                    <div style={{padding: '4px'}}>
                        <CreateProject />
                    </div>
                }
                </Col>
            </Row>
        </div>
    )

}

export default ProjectDetails;