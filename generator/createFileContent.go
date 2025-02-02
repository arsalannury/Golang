package generator

const createFileContent string = `import { ReactNode } from 'react';
import Formik from '@core/component/Formik';
import Paper from '@core/component/Paper';
import BaseMessages from '@core/i18n/BaseMessages';
import BaseView from '@core/view/BaseView';
import Validator from '@core/view/Validator';
import Shortcut from '@core/view/shortcut/Shortcut';
import CheckAuthorization from '@portalclient/security/CheckAuthorization';
import Permission from '../Permission';
import CityModel from './CityModel';
import EmptyModel from '@core/model/EmptyModel';

export default class PageCreateView extends BaseView<any, number> {
  
  // @CheckAuthorization(Permission.ANYTHING)
  componentDidMount(): void {
    super.componentDidMount();
    const id = this.getIdParam();
    if (id) {}
  }

  renderContent(): ReactNode {

    return (
      <Paper>
        <Formik view={this} onSubmit={(model) => this.save(model)}>
          {({ handleSubmit }) => (<></>)}
        </Formik>
      </Paper>
    );
  }

  addValidators(validator: Validator): void {}

  save(model: CityModel): void {}

  getShortcuts(): Shortcut[] {
    return this.getCreateViewShortcuts();
  }

  getViewTitle(): string {
    return this.getMessage(this.getIdParam() ? BaseMessages.EDIT_A : BaseMessages.CREATE_A, { 0: this.getSingularMessage() });
  }

  createModel(props?: Partial<any>): any {
    return new EmptyModel(props);
  }

}

`
