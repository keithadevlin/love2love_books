resource "aws_s3_bucket" "sales_report_download_s3" {
  bucket = "sales-report-download-${var.environment}"
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "get_mws_sales_report_lambda" {
  filename         = "../../../../bin/get_mws_sales_report.zip"
  function_name    = "get-mws-sales-report-${var.environment}"
  role             = "${aws_iam_role.iam_for_lambda.arn}"
  handler          = "get_mws_sales_report"
  source_code_hash = "${filebase64sha256("../../../../bin/get_mws_sales_report.zip")}"
  runtime          = "go1.x"
  timeout          = "900"

  environment {
    variables = {
      SALES_REPORT_BUCKET_NAME = "${aws_s3_bucket.sales_report_download_s3.bucket}"
      REGION                   = "eu-west-1"
    }
  }
}

resource "aws_cloudwatch_event_rule" "every_ten_minute" {
  name                = "request_mws_sales_report"
  description         = "Request MWS Sales Report for last 24 hours"
  schedule_expression = "rate(2 minutes)"

  #schedule_expression = "cron(30 * * * ? *)"
  is_enabled = "false"

  depends_on = [
    "aws_lambda_function.get_mws_sales_report_lambda",
  ]
}

resource "aws_cloudwatch_event_target" "check_every_ten_minute" {
  rule      = "${aws_cloudwatch_event_rule.every_ten_minute.name}"
  target_id = "lambda"
  arn       = "${aws_lambda_function.get_mws_sales_report_lambda.arn}"
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_get_mws_sales_report_lambda" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = "get-mws-sales-report-${var.environment}"
  principal     = "events.amazonaws.com"
  source_arn    = "${aws_cloudwatch_event_rule.every_ten_minute.arn}"
}

# add the lambda policy for logging and access to s3
resource "aws_iam_policy" "lambda_policy" {
  name        = "lambda_policy"
  path        = "/"
  description = "IAM policy for for lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    },
  {
        "Effect": "Allow",
        "Action": [
            "s3:*"
        ],
        "Resource": "arn:aws:s3:::*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = "${aws_iam_role.iam_for_lambda.name}"
  policy_arn = "${aws_iam_policy.lambda_policy.arn}"
}
