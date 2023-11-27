import SubProjects from '../../components/subProjects/SubProjects';
import { Row, Col, Button } from 'antd';
import { useSelector } from 'react-redux';
import { State } from '../../interfaces/state';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { useNavigate } from 'react-router-dom';

const ProjectDetails  = () => {
    const navigate = useNavigate();
    const userPrivileges = useSelector((state : State) => state.user.privileges)
    return (
        <Row>
            <Col span={24}>
                <div style={{display: 'flex', justifyContent: 'flex-end', paddingBottom: '8px'}}>
                    <Button onClick={() => navigate("/create-sub-project")}>Create new sub project</Button>
                </div>
            </Col>
            <Col span={24}>
            {hasPrivilege(userPrivileges, PRIVILEGES.sub_project_read) && <SubProjects />}   
            </Col>
        </Row>
    )
}

export default ProjectDetails;