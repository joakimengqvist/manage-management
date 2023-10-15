import SubProjects from '../components/subProjects/SubProjects';
import { Row, Col, Button } from 'antd';
import { useSelector } from 'react-redux';
import { State } from '../types/state';
import { hasPrivilege } from '../helpers/hasPrivileges';
import { PRIVILEGES } from '../enums/privileges';
import { useNavigate } from 'react-router-dom';

const ProjectDetails: React.FC = () => {
    const navigate = useNavigate();
    const userPrivileges = useSelector((state : State) => state.user.privileges)
    return (
        <div style={{padding: '12px 8px'}}>
            <Row>
                <Col span={24}>
                    <div style={{display: 'flex', justifyContent: 'flex-end', paddingBottom: '8px', paddingRight: '4px'}}>
                        <Button type="primary" onClick={() => navigate("/create-sub-project")}>Create new sub project</Button>
                    </div>
                </Col>
                <Col span={24}>
                {hasPrivilege(userPrivileges, PRIVILEGES.sub_project_read) &&
                    <div style={{padding: '4px'}}>
                        <SubProjects />
                    </div>
                }   
                </Col>
            </Row>
        </div>
    )

}

export default ProjectDetails;