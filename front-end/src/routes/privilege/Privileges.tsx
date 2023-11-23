import Privileges from '../../components/privileges/Privileges';
import CreatePrivilege from '../../components/privileges/CreatePrivilege';
import { Row, Col } from 'antd';
import { useSelector } from 'react-redux';
import { State } from '../../types/state';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';

const PrivilegesDetails: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)
    return (
        <Row>
            <Col span={16}>
                <div style={{paddingRight: '8px'}}>
                    {hasPrivilege(userPrivileges, PRIVILEGES.privilege_read) && <Privileges />}
                </div>
            </Col>
            <Col span={8}>
            {hasPrivilege(userPrivileges, PRIVILEGES.privilege_write) && <CreatePrivilege /> }
            </Col>
        </Row>
    )

}

export default PrivilegesDetails;