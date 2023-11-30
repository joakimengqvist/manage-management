import SubProjects from '../../components/subProjects/SubProjects';
import { Row, Col, Button } from 'antd';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { useNavigate } from 'react-router-dom';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const ProjectDetails  = () => {
    const navigate = useNavigate();
    const userPrivileges = useGetLoggedInUserPrivileges();
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