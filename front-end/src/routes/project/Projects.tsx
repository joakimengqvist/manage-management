import Projects from '../../components/projects/Projects';
import { Row, Col, Button } from 'antd';
import { useSelector } from 'react-redux';
import { State } from '../../types/state';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { useNavigate } from 'react-router-dom';

const ProjectDetails: React.FC = () => {
    const navigate = useNavigate();
    const userPrivileges = useSelector((state : State) => state.user.privileges)
    return (
        <Row>
            <Col span={24}>
                <div style={{display: 'flex', justifyContent: 'flex-end', paddingBottom: '8px'}}>
                    <Button onClick={() => navigate("/create-project")}>Create new project</Button>
                </div>
            </Col>
            <Col span={24}>
            {hasPrivilege(userPrivileges, PRIVILEGES.project_read) && <Projects />}
            </Col>
        </Row>
    )

}

export default ProjectDetails;