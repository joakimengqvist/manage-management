import Privilege from '../../components/privileges/Privilege';
import { Row, Col } from 'antd';
import { useSelector } from 'react-redux';
import { State } from '../../types/state';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';

const PrivilegeDetails: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)
    return (
        <Row>
            <Col span={8}>
            {hasPrivilege(userPrivileges, PRIVILEGES.privilege_read) &&
                <Privilege />
            }
            </Col>
            <Col span={16}>
            </Col>
        </Row>
    )

}

export default PrivilegeDetails;