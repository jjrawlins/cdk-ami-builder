// ~~ Generated by projen. To modify, edit .projenrc.ts and run "npx projen".
import * as path from 'path';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import { Construct } from 'constructs';

/**
 * Props for CheckStateMachineStatusFunction
 */
export interface CheckStateMachineStatusFunctionProps extends lambda.FunctionOptions {
}

/**
 * An AWS Lambda function which executes src/Lambdas/CheckStateMachineStatus/CheckStateMachineStatus.
 */
export class CheckStateMachineStatusFunction extends lambda.Function {
  constructor(scope: Construct, id: string, props?: CheckStateMachineStatusFunctionProps) {
    super(scope, id, {
      description: 'src/Lambdas/CheckStateMachineStatus/CheckStateMachineStatus.lambda.ts',
      ...props,
      runtime: new lambda.Runtime('nodejs20.x', lambda.RuntimeFamily.NODEJS),
      handler: 'index.handler',
      code: lambda.Code.fromAsset(path.join(__dirname, '../../../assets/Lambdas/CheckStateMachineStatus/CheckStateMachineStatus.lambda')),
    });
    this.addEnvironment('AWS_NODEJS_CONNECTION_REUSE_ENABLED', '1', { removeInEdge: true });
  }
}